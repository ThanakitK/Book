<template>
    <v-app-bar elevation="4" color="primary" app clipped-left>
        <v-app-bar-nav-icon class="text1--text" @click="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title class="text1--text">{{ title }}</v-toolbar-title>
        <v-spacer></v-spacer>
        <div class="pt-6">
            <v-switch color="secondary" inset v-model="$vuetify.theme.dark"></v-switch>
        </div>
        <v-btn color="secondary" v-on:click = "$emit('logout','')">Log out</v-btn>
        
    </v-app-bar>
</template>

<script>
export default {
    props: {
        title: {
            type: String,
            default: 'Bank'
        },
        value:{
            type: Boolean,
            default: false
        }
    },
    computed: {
        drawer: {
            set(val){
                this.$emit('input', val)
            },
            get(){
                return this.value
            }
        }
    },
    
}
</script>

<style lang="scss" scoped>

</style>