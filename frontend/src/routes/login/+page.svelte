<script>
    import { GetQRCode, UserInfo } from "$lib/wailsjs/go/main/App";
    import { EventsEmit, EventsOn } from "$lib/wailsjs/runtime/runtime";
    import { onMount } from "svelte";
    import QRCode from "qrcode-svg";
    import { redirect } from "@sveltejs/kit";
    import { goto } from "$app/navigation";
    let svg = $state("")
    let status = $state(0)
    const getQRCode = async () => {
        status = 0
        const res = await GetQRCode()
        console.log(res)
        const qr = new QRCode({
            content: res.data.url,
            padding: 4,
            width: 256,
            height: 256,
            color: "#fb7299",
            background: "#ffffff",
            ecl: "M",
        });
        svg = qr.svg();
        EventsOn(res.data.qrcode_key + "-status", (data) => {
            status = data
            if (status === 3) {
                goto("/")
            }
        });
    }

    onMount(() => {
        UserInfo().then((response) => {
            goto("/")
        }).catch((error) => {
            console.error("Error fetching user info:", error);
            getQRCode()
        });
    });
</script>

<div class="flex w-full h-screen items-center justify-center flex-col">
    {#if svg}
        <h1 class="qrcode-title">扫码登录</h1>
        <div 
            class="relative"
        >
            <img class="w-full"
                class:blur-sm={status != 0}
                src={`data:image/svg+xml;base64,${btoa(unescape(encodeURIComponent(svg)))}`} 
                alt="qrcode">
            <button 
                aria-label="refresh" 
                class="absolute top-0 w-full aspect-square" 
                class:cursor-pointer={status != 0}
                onclick={getQRCode}></button>
        </div>
        {#if status == 1 || status == -1}
            <span>点击二维码刷新</span>
        {:else if status == 2}
            <span>请在手机上确认</span>
        {/if}
    {/if}
</div>

<style>
</style>