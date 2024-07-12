<template>
  <div>
    <button class="shopping-button" @click="expandCart()">Cart</button>
    <NavBar :cartItems="cartItems" />
    <MenuComponent @add-to-cart="addToCart" />
    <Cart class="cart-bar" :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" />
  </div>
</template>

<script>
import NavBar from './components/NavBar.vue';
import MenuComponent from './components/MenuComponent.vue';
import Cart from './components/Cart.vue';

export default {
  components: {
    NavBar,
    MenuComponent,
    Cart
  },
  data() {
    return {
      cartItems: [],
      showCart: false
    };
  },
  methods: {
    expandCart(){
      if (this.showCart == false){
        document.getElementById("mySideNav").style.width="250px";
      } else{
        document.getElementById("mySideNav").style.width="0px";
      }
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
  z-index: 100001;
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