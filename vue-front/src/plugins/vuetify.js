import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
    theme:{
        themes: {
            light:{
                primary: '#4fc08d',
                secondary: '#e91e63',
                text1: '#ffffff',
            },
            dark:{
                primary: '#454545',
                secondary: '#e91e63',
                text1: '#ffffff',
            }
        }
    }
});
