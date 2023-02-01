<template>
    <div>
        <h1>hello my fried</h1>
        <NewPlayerForm/>
    </div>



</template>

<script>
import NewPlayerForm from '@/components/NewPlayerForm.vue'
export default {
    components: {
        NewPlayerForm
    },
}
</script>

<style scoped>
h1 {
    color: black;
    text-align: center;
}
</style>
