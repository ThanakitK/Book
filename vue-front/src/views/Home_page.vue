<template>
    <div>
        <toolbar v-model="drawer" title="App Vue.js" @logout="logout"></toolbar>
        <slidebar v-model="drawer" :menu="menu"></slidebar>
        <v-main>
            <v-container class="">
                <router-view></router-view>
            </v-container>
        </v-main>
    </div>
</template>

<script>
import toolbar from '@/components/layout/com-toolbar.vue'
import slidebar from '@/components/layout/com-slidebar.vue'
export default {
    components: {
        toolbar,
        slidebar
    },
    data(){
        return {
            drawer: false,
            menu: [
                {name: 'page1', path: 'page1', icon: 'mdi-account-hard-hat', text:'Page 1'},
                {name: 'page2', path: 'page2', icon: 'mdi-account-hard-hat', text:'Page 2'}
            ],
                
        }
    },
    methods:{
        logout(){
            localStorage.removeItem('token')
            this.$router.push({name: 'login'})
        }
    }
}
</script>

<style lang="scss" scoped>

</style>