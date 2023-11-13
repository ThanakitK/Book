<template>
    <div>
       
        <h1>Add Book</h1>
        <v-text-field label="Title" v-model="payload.title"></v-text-field>
        <v-text-field label="Price" v-model="payload.price" type="number"></v-text-field>
        <v-btn @click="add">Add</v-btn>
        <v-snackbar  v-model="snackbar" color='success' :timeout="2000">
            <v-icon>mdi-checkbox-marked-circle</v-icon>
            Add book success !!!
        </v-snackbar>
    </div>
</template>

<script>
export default {
    data(){
        return {
            snackbar: false,
            payload: {
                title: '',
                price: null
            },
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
                    this.payload.title = '';
                    this.payload.price = null;
                    this.snackbar = true;
                    console.log(res);
                })
                .catch((err) => {
                    console.log(err);
                });

            
        }, 
        
        clearSuccessMessage() {
            this.successMessage = '';
        },
    }
}
</script>

<style lang="scss" scoped>

</style>