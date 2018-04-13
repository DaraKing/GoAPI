var eventHub = new Vue();
const API_URL = 'http://localhost:8081';

Vue.component('modal', {
    data: function(){
        return {
            message: '',
            id: '',
            firstName: '',
            lastName: '',
            age: '',
            email: '',
            gender: '',
            userDefined: false,
        }
    },
    created: function () {
        eventHub.$on('getData', this.getData);
    },
    destroyed: function () {
        eventHub.$off('getData', this.getData);
    },
    template: '#modal-template',
    methods: {
        getData: function(data){
            this.id = data.id,
            this.firstName = data.first_name,
            this.lastName = data.last_name,
            this.age = data.age,
            this.email = data.email,
            this.gender = data.gender,
            this.userDefined = true
        },
        editUser: function () {

            var JSON =
                {
                    first_name: this.firstName, last_name: this.lastName,
                    age: parseInt(this.age), email: this.email, gender: this.gender
                };

            this.$http.post(API_URL+'/edit/' + this.id + '/', JSON).then(response => {
                if (response.status === 200) {
                    this.message = 'Successfuly edited!';
                    eventHub.$emit('refreshTable');
                }
            });
        },
        addUser: function () {
            var JSON =
                {
                    first_name: this.firstName, last_name: this.lastName,
                    age: parseInt(this.age), email: this.email, gender: this.gender
                };
            this.$http.post(API_URL+'/insertUser', JSON).then(response => {
                if (response.status === 200) {
                    this.message = 'Successfuly added!';
                    eventHub.$emit('refreshTable');
                }
            });
        }
    }
});

const app = new Vue({
    el: '#app',
    data: {
        results: null,
        search: '',
        showModal: false,
    },
    created: function () {
        eventHub.$on('refreshTable', this.refreshTable);
        this.getUsers();
    },
    methods: {
        getUsers: function () {
            this.$http.get(API_URL+'/all').then(response => {
                this.results = response.body;
            });
        },
        deleteUser: function (event) {
            var id = event.currentTarget.id;

            this.$http.delete(API_URL+'/delete/' + id + '/').then(response => {
                if (response.status === 200) {
                    this.getUsers();
                }
            });
        },
        showModalFunc: function(user) {
            this.showModal = true;
            if(user){
                this.$nextTick(() => {
                    eventHub.$emit('getData', user);
                });
            }
        },
        refreshTable: function() {
            this.getUsers();
        }
    }
});