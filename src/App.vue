<template>
  <div id="app">
    <NavBar @remove-item="removeItem" :show-cart="showCart" @expand-cart="expandCart" :cartItems="cartItems"  @clear-cart="clearCart" />
    <div class="mt-12">
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
      cartItems: [],
      showCart: false
    };
  },
  methods: {
    expandCart(){
      this.showCart=!this.showCart    
    },
    addToCart(item) {
      this.cartItems.push(item); // Add item to cart
      this.saveCart(); // Save cart state to localStorage
    },
    removeItem(index) {
      this.cartItems.splice(index, 1); // Remove item from cart
      this.saveCart(); // Save cart state to localStorage
    },
    clearCart() {
      this.cartItems = []; // Clear cartItems array
      this.saveCart(); // Save empty cart state to localStorage
    },
    saveCart() {
      localStorage.setItem('cartItems', JSON.stringify(this.cartItems));
    },
    loadCart() {
      const cartItems = localStorage.getItem('cartItems');
      if (cartItems) {
        this.cartItems = JSON.parse(cartItems);
      }
    }
  },
  created() {
    this.loadCart();
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
