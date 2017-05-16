<template>
  <div class="hello">
    <!--<img src="../assets/logo.png">-->
    <md-toolbar>
      <md-button class="md-icon-button" @click.native="toggleSideNav">
        <md-icon>menu</md-icon>
      </md-button>

      <h2 class="md-title">Editor</h2>
    </md-toolbar>

    <md-sidenav class="md-left" ref="sideNav">
      <md-toolbar class="md-large">
        <div class="md-toolbar-container">
          <h3 class="md-title">Sidenav content</h3>
        </div>
      </md-toolbar>

      <p>Linkz...</p>
    </md-sidenav>

    <div class="main-content">
      <p>{{ msg }}</p>
    </div>

  </div>
</template>

<script>
export default {
  name: 'main',
  data () {
    return {
      msg: 'Hellou'
    }
  },
  methods: {
    toggleSideNav () {
      this.$refs.sideNav.toggle()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
