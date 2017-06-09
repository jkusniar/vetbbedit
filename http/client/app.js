"use strict";

Vue.use(VueMaterial);

// vm is vue instance and error is error result returned by axios
function showError(vm, error) {
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
}

new Vue({
    el: '#app',
    data: {
        news: [],
        newNews: "",
        services: [],
        newService: "",
        hours: {},
        newFootnote: "",
        respMsg: "",
        prompt: {
            title: 'Set hours',
            placeholder: 'Type hours...',
            value: "xx",
            ampm: "",
            index: -1
        }
    },
    mounted: function () {
        this.getData();
    },
    methods: {
        getData: function () {
            this.news = [];
            var vm = this;
            axios.get('/data')
                .then(function (response) {
                    vm.news = response.data.news.items;
                    vm.services = response.data.services.items;
                    vm.hours = response.data.hours;
                })
                .catch(function (error) {
                    showError(vm, error);
                });
        },
        onNewNews: function (event) {
            if (this.newNews) {
                this.news.push(this.newNews);
                this.newNews = "";
            }
        },
        removeNews: function (index) {
            this.news.splice(index, 1);
        },
        onNewService: function (event) {
            if (this.newService) {
                this.services.push(this.newService);
                this.newService = "";
            }
        },
        removeService: function (index) {
            this.services.splice(index, 1);
        },
        onNewFootnote: function (event) {
            if (this.newFootnote) {
                this.hours.footnotes.push(this.newFootnote);
                this.newFootnote = "";
            }
        },
        removeFootnote: function (index) {
            this.hours.footnotes.splice(index, 1);
        },
        showEditDateDialog: function (index) {
            console.log(index);
        },
        openDialog: function (ref, ampm, index) {
            // set md-open-from="#id-of-icon-component" md-close-to="#id-of-icon-component"
            this.prompt.ampm = ampm;
            this.prompt.index = index;

            if (ampm === "am") {
                this.prompt.value = this.hours.days[index].am;
            } else if (ampm === "pm") {
                this.prompt.value = this.hours.days[index].pm;
            } else {
                console.log("Bad value for ampm: " + ampm);
            }

            this.$refs[ref].open();
        },
        closeDialog: function (ref) {
            this.$refs[ref].close();
        },
        onInput: function (newValue) {
            if (this.prompt.ampm === "am") {
                this.hours.days[this.prompt.index].am = newValue;
            } else if (this.prompt.ampm === "pm") {
                this.hours.days[this.prompt.index].pm = newValue;
            } else {
                console.log("Bad value for ampm: " + this.prompt.ampm);
            }
        },
        saveAndUpload: function () {
            this.respMsg = "";
            var vm = this;
            axios.put('/data', {
                news: {items: this.news},
                services: {items: this.services},
                hours: this.hours
            })
                .then(function (response) {
                    vm.respMsg = "save OK";
                    vm.showMessage()
                })
                .catch(function (error) {
                    showError(vm, error);
                });
        },
        showMessage: function () {
            this.$refs.snackbar.open();
        }
    }
});