Vue.config.delimiters= ['((', '))'];

const Foo = { template: '<div>foo</div>' }
const Bar = { template: '<div>bar</div>' }

const routes = [
    { path: '/foo', component: Foo },
    { path: '/bar', component: Bar }
];

const router = new VueRouter({
    routes // short for `routes: routes`
});

window.onload = function(){
    var app = new Vue ({
        el: '#app',
        router,
        data: {
            results: null,
            search: '',
            editMode: false,
            editedUser: null,
            firstName: '',
            lastName: '',
            age: '',
            email: '',
            gender: '',
        },
        created: function(){
            this.getPhones();
        },
        methods: {
            getPhones: function() {
                    this.$http.get('http://localhost:8081/all').then(response => {

                        // get body data
                        //console.log(response.body);
                        this.results = response.body;
                        console.log(response.status);

                }, response => {
                        // error callback
                    console.log(response);
                    });
                },
            submit: function () {
                var JSON =
                    {first_name: this.firstName, last_name: this.lastName,
                        age: this.age, email: this.email, gender: this.gender
                    };
                console.log(JSON);
                this.$http.post('http://localhost:8081/insertUser', JSON).then(response => {
                    console.log(response.status);
                });
            },
            onClick: function (event) {
                var id=event.currentTarget.id;

                this.$http.delete('http://localhost:8081/delete/'+ id+'/').then(response => {
                    console.log(response.status);
                });
            },
            edit: function (event) {
                var elementId=event.currentTarget.id;
                var id = parseInt(elementId.replace(/\D/g, ''));
                var userFirstName = document.getElementById("firstName"+id).value,
                    userLastName = document.getElementById("lastName"+id).value,
                    userAge = parseInt(document.getElementById("age"+id).value),
                    userEmail = document.getElementById("email"+id).value,
                    userGender = document.getElementById("gender"+id).value;

                console.log()

                var JSON = {first_name: userFirstName, last_name: userLastName, age: userAge, email: userEmail, gender: userGender};

                this.$http.post('http://localhost:8081/edit/'+id+'/', JSON).then(response => {
                    console.log(response.status);
                });
            },
        }
    });
}