<template>
  <div class="container">
    <md-sidenav class="main-sidebar md-left md-fixed" md-swipeable ref="main-sidebar">
      <md-toolbar>
        <h1 class="md-title">MARIE</h1>
      </md-toolbar>
  
      <div class="main-sidebar-links">
        <md-list class="md-dense">
          <md-list-item>
            <router-link exact to="/">Things</router-link>
            <div class="md-list-action">
              <md-button class="md-icon-button">
                <router-link to="/thing-form">
                  <md-icon>add</md-icon>
                </router-link>
              </md-button>
            </div>
          </md-list-item>
        </md-list>
      </div>
  
    </md-sidenav>
    <transition name="md-router" appear>
      <router-view></router-view>
    </transition>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style lang="scss">
@import './assets/variables.scss';

$sizebar-size: 280px;


[v-cloak] {
  display: none;
}

html,
body {
  height: 100%;
}

body {
  display: flex;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

.container {
  min-height: 100%;
  display: flex;
  flex-flow: column nowrap;
  flex: 1;
  transition: $swift-ease-out;
  @media (min-width: 1281px) {
    padding-left: $sizebar-size;
  }
}

.main-sidebar.md-sidenav {
  .md-sidenav-content {
    width: $sizebar-size;
    display: flex;
    flex-flow: column;
    overflow: hidden;
    @media (min-width: 1281px) {
      top: 0;
      pointer-events: auto;
      transform: translate3d(0, 0, 0) !important;
      box-shadow: $material-shadow-2dp;
    }
  }
  .md-backdrop {
    @media (min-width: 1281px) {
      opacity: 0;
      pointer-events: none;
    }
  }
  .md-toolbar {
    min-height: 172px;
    border-bottom: 1px solid rgba(#000, .12);
  }
  .vue-material-logo {
    font-size: 24px;
    a {
      width: 100%;
      display: flex;
      flex-flow: column;
      justify-content: center;
      align-items: center;
      color: inherit;
      text-decoration: none;
      &:hover {
        color: inherit;
        text-decoration: none;
      }
    }
    img {
      width: 160px;
      margin-bottom: 16px;
    }
  }
  .main-sidebar-links {
    overflow: auto;
    flex: 1;
    .md-inset .md-list-item-container {
      padding-left: 36px;
    }
    .md-list-item-container {
      font-size: 14px;
      font-weight: 500;
    }
  }
  .release-version {
    padding: 8px 8px 8px 16px;
    border-top: 1px solid rgba(#000, .12);
    display: none;
    @media (max-width: 480px) {
      display: block;
    }
    >div {
      justify-content: center;
    }
    .md-select:after {
      color: rgba(#000, .87);
    }
  }
}

</style>
