<template>
  <div>
    <!-- <router-view /> -->
    <!-- <NavBar @remove-item="removeItem" :show-cart="showCart" @expand-cart="expandCart" :cartItems="cartItems"  @clear-cart="clearCart" /> -->

    <!-- <Cart :show-cart="showCart" id="mySideNav" @click="closePopupOnOutsideClick" class="cart-bar" @expand-cart="expandCart" :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" /> -->
    <!-- <Cart id="mySideNav" @click="closePopupOnOutsideClick" class="cart-bar" :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" /> -->

    <Menu @expand-cart="expandCart" class="md:mt-14 z-0" @add-to-cart="addToCart" />
    <div class=" z-[1000]">{{ showCart }}</div>
  </div>
</template>


<script>
import NavBar from '../components/NavBar.vue';
import Menu from '../components/Menu.vue';

export default {
  components: {
    NavBar,
    Menu
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
