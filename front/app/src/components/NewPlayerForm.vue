<template>
    <div class="formUser">
        <form class="form">
            <h4>Введите</h4>
            <input
            type="text"
            class="input"
            placeholder="nickname"
            >
            <input
            type="text"
            class="input"
            placeholder="money in the account"
            >
            <button class="button" @click="$router.push('/lobbies')">
            Create
            </button>
        </form>
    </div>
</template>

<script>
export default {

}
</script>

<style>

.formUser {
    display: flex;
    justify-content: center;
}

h4 {
    font-family: 'Courier New', Courier, monospace;
    text-align: center;
    font-size: 30px;
}

.form {
    display: flex;
    flex-direction: column;
}

.input {
    background: none;
    border: 3px solid teal;
    width: 200px;
    height: 25px;
    margin: 20px;
}

.input:focus {
    outline: none;
    box-shadow: 0 0 0 2px #46dc87;
    border: 3px solid #46dc87;
}

.button {
    appearance: none;
    border: 0;
    border-radius: 5px;
    background: #1daa9a;
    color: #fff;
    padding: 8px 16px;
    font-size: 16px;
    margin-top: 20px;
}

.button:focus {
    outline: none;
    box-shadow: 0 0 0 4px #cbd6ee;
}

.button:hover {
    background: #46dc87;
}
</style>
