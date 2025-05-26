<script lang="ts">
    import { GetLiveCategory, UserInfo, StartLive, StopLive, UpdateLiveInfo, GetRoomStatus } from "$lib/wailsjs/go/main/App";
    
    import { onDestroy, onMount } from "svelte";
    import { goto } from "$app/navigation";

    import { toast } from "svelte-sonner";
    import * as Select from "$lib/components/ui/select/index"
    import * as Dialog from "$lib/components/ui/dialog/index.js";
    import { Input } from "$lib/components/ui/input";
    import { Button, buttonVariants } from "$lib/components/ui/button";

    let intervalID = null
    
    let user = $state({
        name: "",
        avatar: "",
        roomid: 0,
    })
    let cat = $state({
        data: [],
        code: 0,
        message: "",
        idToName: {}
    })
    let rtmp = $state({
        addr: "",
        code: ""
    })
    let roomStatus = $state({
        code: 0,
        message: "",
        status: 0
    })
    let title = $state("")
    let mainID = $state(-1)
    let subID = $state(-1)
    const startLive = async () => {
        if (title === "" || title.length > 80) {
            toast.error("标题不能为空或超过80个字符")
            return
        }
        try {
            await UpdateLiveInfo(title, subID)
            const response = await StartLive(title, subID);
            rtmp = {
                addr: response.addr,
                code: response.code
            }
            toast.success("已开启直播")
        } catch (error) {
            console.error("Error during startLive:", error);
            toast.error("开启直播失败")
            return;
        }
    }
    const stopLive = async () => {
        try {
            await StopLive()
            rtmp = {
                addr: "",
                code: ""
            }
            toast.success("直播已关闭")
        } catch (error) {
            console.log(error)
            toast.error("关闭直播失败")
        }
    }
    onMount(async () => {
        try {
            const userInfo = await UserInfo()
            user = {
                name: userInfo.name,
                avatar: userInfo.avatar,
                roomid: userInfo.room_id,
            };
            intervalID = setInterval(async () => {
                try {
                    const response = await GetRoomStatus(userInfo.room_id);
                    roomStatus = {
                        code: response.code,
                        message: response.message,
                        status: response.status
                    }
                    title = response.title || "";
                    mainID = response.parent_area_id || -1;
                    subID = response.area_id || -1;
                } catch (error) {
                    console.error("Error fetching room status:", error);
                }
            }, 2000);           
        } catch (error) {
            goto("/login");
        }
        try {
            const response = await GetLiveCategory()
            if (response.code !== 0) {
                toast.error("获取直播分区列表失败，请重新打开软件")
            } else {
                cat = {
                    data: response.data,
                    code: response.code,
                    message: response.message,
                    idToName: {}
                }
                response.data.forEach((item) => {
                    cat.idToName[item.id] = {
                        name: item.name,
                        subs: item.list
                    };
                });
            }
        } catch (error) {
            toast.error("获取直播分区列表失败，请重新打开软件")
            console.error("Error fetching live category:", error);
        }
    })
    onDestroy(() => {
        if (intervalID) {
            clearInterval(intervalID);
        }
    });
</script>

<div class="flex mx-auto flex-1 w-2/3 p-4 overflow-y-auto items-center justify-center flex-col">
    <div class="flex flex-col gap-2 p-8 sm:flex-row sm:items-center sm:gap-6 sm:py-4 ...">
      <img referrerpolicy="no-referrer" class="mx-auto block h-24 rounded-full sm:mx-0 sm:shrink-0" src={user.avatar} alt="" />
      <div class="space-y-2 text-center sm:text-left">
        <div class="space-y-0.5">
          <p class="text-lg font-semibold text-black">{user.name}</p>
          <p class="font-medium text-gray-500">房间号：{user.roomid}</p>
        </div>
      </div>
    </div>
    <div class="flex w-full max-w-sm flex-col gap-1.5">
        <div class="flex max-w-sm flex-col gap-1.5">
            <Input type="text" id="title" placeholder="请输入标题" bind:value={title}/>
        </div>
        <div class="flex flex-col gap-2 sm:flex-row sm:items-center sm:gap-6 sm:py-4 ...">
            <Select.Root type="single" name="main-category" bind:value={mainID}>
                <Select.Trigger class="w-1/2">
                    {mainID === -1 ? "选择分区" : cat.idToName[mainID].name}
                </Select.Trigger>
                <Select.Content>
                    {#each cat.data as item}
                        <Select.Item value={item.id} label={item.name}/>
                    {/each}
                </Select.Content>
            </Select.Root>
            <Select.Root name="sub-category" type="single" bind:value={subID} disabled={mainID === -1}>
                <Select.Trigger class="w-1/2">
                    {subID === -1 ? "选择子分区" : cat.idToName[mainID].subs.find(item => item.id === subID)?.name}
                </Select.Trigger>
                <Select.Content>
                    {#each cat.idToName[mainID].subs as item}
                        <Select.Item value={item.id} label={item.name}/>
                    {/each}
                </Select.Content>
            </Select.Root>
        </div>
        {#if rtmp.addr !== "" }
            <Dialog.Root>
                <Dialog.Trigger class={buttonVariants({ variant: "outline" })}>
                    查看推流码
                </Dialog.Trigger>
                <Dialog.Content>
                  <Dialog.Header>
                      <Dialog.Title>推流码</Dialog.Title>
                      <Dialog.Description>
                          <span>RTMP服务器地址</span>
                          <p class="break-normal md:break-all">{rtmp.addr}</p>
                          <br>
                          <span>推流码</span>
                          <p class="break-normal md:break-all">{rtmp.code}</p>
                      </Dialog.Description>
                  </Dialog.Header>
                </Dialog.Content>
            </Dialog.Root>
            <Button onclick={stopLive}>关闭直播</Button>
        {:else if rtmp.addr === "" && roomStatus.status === 1}
            <Button onclick={startLive}>重新开播</Button>
            <Button onclick={stopLive}>关闭直播</Button>
        {:else}
            <Button onclick={startLive}>开始直播</Button>
        {/if}
    </div>
</div>
