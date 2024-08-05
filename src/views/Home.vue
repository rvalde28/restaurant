<template>
  <div>
    <Menu @expand-cart="expandCart" class="mt-14 z-0" />
    <div class=" z-[1000]">{{ showCart }}</div>
  </div>
</template>


<script>
import Menu from '../components/Menu.vue';

export default {
  components: {
    Menu
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
  }
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
