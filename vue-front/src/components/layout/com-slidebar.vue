<template>
    <v-navigation-drawer v-model="drawer" app clipped buttom>
        <v-list nav dense>
            <v-list-item-group active-class="primary" class="text1--text">
                <v-list-item v-for="item in menu" :key="item.name" :to="item.path">
                    <v-list-item-icon>
                        <v-icon>{{ item.icon }}</v-icon>
                    </v-list-item-icon>
                    <v-list-item-title>
                        {{ item.text }}
                    </v-list-item-title>
                </v-list-item>
            </v-list-item-group>
        </v-list>
    </v-navigation-drawer>
</template>

<script>
export default {
    props: {
        menu:{
            type: Array,
            default: ()=> [],
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