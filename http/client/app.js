"use strict";

Vue.use(VueMaterial);

new Vue({
    el: '#app',
    data: {
        greeting: 'Hellou!'
    },
    methods: {
        toggleSideNav: function() {
            this.$refs.sideNav.toggle();
        }
    }
});