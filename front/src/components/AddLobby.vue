<template>
    <div class="dialog" v-if="show">
        <div class="dialog__content">
            <button class="button" @click="hideDialog">X</button>
            <p class="text">Name lobby</p>
            <input class="input" type="text">
            <p class="text">Max players</p>
            <input class="input" type="text">
            <button class="button" @click="hideDialog">create</button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'my-dialog',
    props: {
        show: {
            type: Boolean,
            default: false
        }
    },
    methods: {
        hideDialog() {
            this.$emit('update:show', false)
        }
    }
}

</script>

<style scoped>

*{
    font-family: Areal;
}
.dialog {
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0, 0, 0, 0.5);
    position: fixed;
    display: flex;
}

.dialog__content {
    margin: auto;
    background: white;
    border-radius: 12px;
    min-height: 70px;
    min-width: 300px;
    padding: 20px;
}

</style>
