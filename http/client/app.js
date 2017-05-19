"use strict";

Vue.use(VueMaterial);

new Vue({
    el: '#app',
    data: {
        news: [],
        newNews: ""
    },
    mounted: function () {
        this.getNews();
    },
    methods: {
        toggleSideNav: function () {
            this.$refs.sideNav.toggle();
        },
        getNews: function () {
            this.news = [];
            var vm = this;
            axios.get('/news')
                .then(function (response) {
                    console.log(response.data);
                    console.log(response.status);
                    console.log(response.statusText);
                    vm.news = response.data;
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        onNewNews: function (event) {
            console.log(event.key);
            if(event.key === "Enter") {
                console.log('the id of the input was:' + event.currentTarget.id);
                console.log(this.newNews);
            }

        }
    }
});