<template>
    <div>
        <v-data-table v-bind:headers="headers" :items="dataApi"></v-data-table>
    </div>
</template>

<script>
export default {
    data(){
        return {
            dataApi: [],
            headers: [{
                text: 'Title',
                align: 'start',
                sortable: false,
                value: 'title'
                },
                {
                text: 'Price',
                align: 'start',
                sortable: false,
                value: 'price'
                }
            ]
        }
    },
    created(){
        const options={
            url:'http://localhost:3000/books',
            method:'get',
        }

        this.axios(options)
            .then((res) => {
                this.dataApi = res.data;
            })
            .catch((err) => {
                console.log(err);
            });
    }
}
</script>

<style lang="scss" scoped>

</style>