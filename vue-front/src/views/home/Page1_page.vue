    <template>
        <div>
            <v-data-table v-bind:headers="headers" :items="dataApi">
                <template v-slot:item="row">
                    <tr>
                        <td v-if="!row.isEditing">{{row.item.title}}</td>
                        <td v-if="!row.isEditing">{{row.item.price}}</td>

                        <!-- <td v-if="row.isEditing"><v-text-field v-model="row.item.title" label="Title"></v-text-field></td>
                        <td v-if="row.isEditing"><v-text-field v-model="row.item.price" label="Price" type="number"></v-text-field></td> -->
                        
                        <td v-if="!row.isEditing">
                            <v-btn @click="edit_book(row.item)">edit</v-btn>
                            <v-btn @click="delete_book(row.item)">delete</v-btn>
                        </td>
                        <!-- <td v-if="row.isEditing">
                            <v-btn @click="edit_book(row.item)">save</v-btn>
                            <v-btn @click="toggleEdit(row.edit_mode)">cancel</v-btn>
                        </td> -->
                        
                    </tr>
                </template>
            </v-data-table>
        </div>
    </template>

    <script>
    export default {
        data(){
            return {
                edit_mode: [],
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
                    },
                    {
                        text: '',
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
        },
        methods: {
            edit_book(item) {
                console.log(item.id);
            },
            delete_book(item) {
                const options={
                url:'http://localhost:3000/delete/'+item.id,
                method:'delete',
                }
                this.axios(options)
                .then((res) => {
                    this.dataApi = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
                window.location.reload();
            },

        }
    }
    </script>

    <style lang="scss" scoped>

    </style>