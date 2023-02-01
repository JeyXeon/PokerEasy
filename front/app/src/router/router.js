import Main from "@/pages/Main"
import Lobbies from '@/pages/Lobbies'
import GameLobby from '@/pages/GameLobby'
import { createRouter, createWebHistory } from "vue-router"

const routes = [
    {
        path: '/',
        component: Main
    },
    {
        path: '/lobbies',
        component: Lobbies
    },
    {
        path: '/lobbies/:id',
        component: GameLobby
    },
]

const router = createRouter({
    routes,
    history: createWebHistory()
})

export default router;
