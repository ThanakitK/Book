<template>
    <div class="center">
        <v-card outlined class="pa-5">
            <v-card-title>{{title}}</v-card-title>
            <v-card-text>
                <v-text-field dense type="" outline color="primary"></v-text-field>
                <v-text-field dense type="" outline color="primary"></v-text-field>
                <v-btn block small elevation="" color="primary" @click="login">Login</v-btn>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
export default {
    props: {
        title: {
            type: String,
            default: 'Login'
        },
    },
    methods:{
        login(){
            localStorage.setItem('token', 'vue.js')
            this.$router.push({name: 'home'});
        }
    }
}
</script>

<style lang="scss" scoped>
    .center{
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
    }
</style>