<template>
  <div id="app">
    <NavBar @remove-item="removeItem" :show-cart="showCart" @expand-cart="expandCart" />
    <div class="md:mt-12">
      <router-view></router-view>
    </div>
  </div>
</template>



<script>
import NavBar from './components/NavBar.vue';

export default {
  name: "app",
  components: {
    NavBar,
  },
  data() {
    return {
      showCart: false
    };
  },
  methods: {
    expandCart(){
      this.showCart=!this.showCart    
    }
  },
};
</script>

<style>
.cart-bar{
  z-index: 100;
}
/* .menu-app{
  overflow:scroll;
  height:auto;
  -webkit-overflow-scrolling: touch;
} */
.shopping-button {
  position: fixed;
  bottom: 12px;
  right: 0;
  border-radius: 50%;
  display: flex;
  width: 60px;
  height: 60px;
}
</style>
