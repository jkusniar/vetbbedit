"use strict";

Vue.use(VueMaterial);

new Vue({
    el: '#app',
    data: {
        news: [],
        newNews: "",
        respMsg: ""
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
                    vm.news = response.data.items;
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        onNewNews: function (event) {
            if (this.newNews) {
                this.news.push(this.newNews);
                this.newNews = "";
            }
        },
        removeItem: function (index) {
            //console.log("deleting " + index)
            this.news.splice(index, 1);
        },
        saveAndUpload: function () {
            this.respMsg = "";
            var vm = this;
            axios.put('/news', {items: this.news})
                .then(function (response) {
                    vm.respMsg = "save OK";
                    vm.showMessage()
                })
                .catch(function (error) {
                    if (error.response) {
                        console.log(error.response.data);
                        console.log(error.response.status);
                        console.log(error.response.headers);
                        vm.respMsg = error.response.data;
                    } else if (error.request) {
                        console.log(error.request);
                        vm.respMsg = "no response";
                    } else {
                        console.log('Error', error.message);
                        vm.respMsg = error.message;
                    }
                    console.log(error.config);

                    vm.showMessage()
                });
        },
        showMessage: function () {
            this.$refs.snackbar.open();
        }
    }
});