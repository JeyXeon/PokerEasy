<template>
    <div class="content">
        <table class="table">
            <tr>
                <th>id</th>
                <th>name</th>
                <th>players</th>
                <th>
                    <button class="button"
                        @click="$router.push(`/lobbies/${lobby.id}`)"  
                    >
                    connect</button>
                </th>
            </tr>
        </table>
    </div>
</template>


<script>
export default {

}
// @click="$router.push(`/lobbies/${lobby.id}`)"  должен быть список лобби, и по нажатию на 'connect' перекидывать на страниицу лобби с его уник id
</script>

<style scoped>


.table {
    background-color: skyblue;
}

.content {
    display: flex;
    flex-direction: column;
    margin: auto;
}
</style>
