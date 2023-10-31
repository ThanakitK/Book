<template>
    <div>
        <h1>Add Book</h1>
        <v-text-field label="Title" v-model="payload.title"></v-text-field>
        <v-text-field label="Price" v-model="payload.price" type="number"></v-text-field>
        <v-btn @click="add">Add</v-btn>
    </div>
</template>

<script>
export default {
    data(){
        return {
            payload: {
                title: '',
                price: null
            }
        }
    },
    methods: {
        async add(){
            const price = parseInt(this.payload.price, 10);
            this.payload.price = price;
            const options={
                url:'http://localhost:3000/book',
                method:'post',
                data: this.payload
            }
            this.axios(options)
                .then((res) => {
                    console.log(res);
                })
                .catch((err) => {
                    console.log(err);
                });

            this.payload.title = '';
            this.payload.price = null;
        }
    }
}
</script>

<style lang="scss" scoped>

</style>