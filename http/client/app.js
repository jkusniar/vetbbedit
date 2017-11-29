"use strict";

Vue.use(VueMaterial);
Vue.use(VueI18n);

var i18n = new VueI18n({
    locale: 'sk',  // TODO: set locale from client
    messages: {
        en: {
            message: {
                news: 'News',
                newsText: 'News text',
                newsPlaceholder: 'Type news and press enter',
                services: 'Services',
                serviceText: 'Service text',
                servicePlaceholder: 'Type service and press enter',
                hours: 'Opening hours',
                day: 'Day',
                am: 'AM',
                pm: 'PM',
                footnoteText: 'New footnote text',
                footnotePlaceholder: 'Type footnote and press enter',
                closeBtn: 'Close',
                savedOk: 'Saved OK',
                errNoResponse: 'ERROR: No response'
            }
        },
        sk: {
            message: {
                news: 'Oznamy',
                newsText: 'Text oznamu',
                newsPlaceholder: 'Napíš oznam a stlač enter',
                services: 'Služby',
                serviceText: 'Znenie služby',
                servicePlaceholder: 'Napíš službu a stlač enter',
                hours: 'Ordinačné hodiny',
                day: 'Deň',
                am: 'Doobedu',
                pm: 'Poobede',
                footnoteText: 'Text poznámky',
                footnotePlaceholder: 'Napíš poznámku a stlač enter',
                closeBtn: 'Zatvoriť',
                savedOk: 'Uložené OK',
                errNoResponse: 'CHYBA: server neodpovedal'
            }
        }
    }
});


// vm is vue instance and error is error result returned by axios
function showError(vm, error) {
    if (error.response) {
        console.log(error.response.data);
        console.log(error.response.status);
        console.log(error.response.headers);
        vm.respMsg = error.response.data;
    } else if (error.request) {
        console.log(error.request);
        vm.respMsg = vm.$t("message.errNoResponse");
    } else {
        console.log('Error', error.message);
        vm.respMsg = error.message;
    }
    console.log(error.config);

    vm.showMessage()
}

new Vue({
    el: '#app',
    i18n: i18n,
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
        },
        uploading: false
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
            vm.uploading = true;
            axios.put('/data', {
                news: {items: this.news},
                services: {items: this.services},
                hours: this.hours
            })
                .then(function (response) {
                    vm.uploading = false;
                    vm.respMsg =  vm.$t("message.savedOk");
                    vm.showMessage()
                })
                .catch(function (error) {
                    vm.uploading = false;
                    showError(vm, error);
                });
        },
        showMessage: function () {
            this.$refs.snackbar.open();
        }
    }
});