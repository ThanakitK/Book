    <template>
        <div>
            <v-data-table v-bind:headers="headers" :items="dataApi">
                <template v-slot:item="row">
                    <tr>
                        <td v-if="editingItemId !== row.item.id">
                            {{row.item.title}}
                            
                        </td>
                        <td v-if="editingItemId == row.item.id">
                            <v-text-field v-model="row.item.title" />
                        </td>
                        <td v-if="editingItemId !== row.item.id">
                            {{row.item.price}}
                        </td>
                        <td v-if="editingItemId == row.item.id">
                            <v-text-field v-model="row.item.price" />
                        </td>

                        
                        <td v-if="editingItemId !== row.item.id">
                            <v-btn @click="editingItemId = row.item.id">edit</v-btn>
                            <v-btn @click="delete_book(row.item)">delete</v-btn>
                        </td>
                        <td v-if="editingItemId == row.item.id">
                            <v-btn @click="edit_book(row.item)">save</v-btn>
                            <v-btn @click="editingItemId = null">cancel</v-btn>
                        </td>
                        
                    </tr>
                </template>
            </v-data-table>
            <v-snackbar  v-model="snackbar" color='success' :timeout="2000">
                <v-icon>mdi-checkbox-marked-circle</v-icon>
                    {{ successMessage }}
            </v-snackbar>
        </div>
    </template>

    <script>
    export default {
        data(){
            return {
                snackbar: false,
                successMessage: '',
                editingItemId: null,
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
                ],
                payload: {
                    title: '',
                    price: null
                }
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
            async edit_book(item) {
                const price = parseInt(item.price, 10);
                this.payload.title = item.title
                this.payload.price = price
                const options={
                url:'http://localhost:3000/update/'+item.id,
                method:'put',
                data: this.payload
                }
                await this.axios(options)
                .then(() => {
                    this.snackbar = true;
                    this.successMessage = 'Book updated success !!!';
                    const fetchData={
                        url:'http://localhost:3000/books',
                        method:'get',
                    }

                    return this.axios(fetchData);
                })
                .then((res) => {
                    // Update the dataApi with the fetched data
                    this.dataApi = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
                
                this.editingItemId = null;
            },
            async delete_book(item) {
                const options={
                url:'http://localhost:3000/delete/'+item.id,
                method:'delete',
                }
                await this.axios(options)
                .then(() => {
                    this.snackbar = true;
                    this.successMessage = 'Book deleted success !!!';
                    const fetchData={
                        url:'http://localhost:3000/books',
                        method:'get',
                    }

                    return this.axios(fetchData);
                })
                .then((res) => {
                    this.dataApi = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
            },

        }
    }
    </script>

    <style lang="scss" scoped>

    </style>