import { writable } from "svelte/store";

const user = writable({
    name: "",
    roomid: 0,
    avatar: ""
})

export {user};