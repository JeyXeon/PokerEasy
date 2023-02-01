<template>
    <div>
        <CreateLobbyButton
        @click="showDialog"
        />
        <RefreshButton/>
        <Table/>
        <AddLobby v-model:show="dialogVisible"/>
    </div>
</template>

<script>
import Table from '@/components/Table'
import CreateLobbyButton from '@/components/CreateLobbyButton'
import RefreshButton from '@/components/RefreshButton'
import AddLobby from '@/components/AddLobby'
export default {
    components: {
        Table, CreateLobbyButton, RefreshButton, AddLobby
    },
    data () {
        return {
            dialogVisible: false,
        }
    },
    methods: {
        showDialog() {
            this.dialogVisible = true;
        }
    }
}

</script>

<style scoped>

</style>
