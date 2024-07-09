<template>
  <div>
    <NavBar :cartItems="cartItems" />
    <Menu @add-to-cart="addToCart" />
    <Cart :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" />
  </div>
</template>

<script>
import NavBar from './components/NavBar.vue';
import Menu from './components/Menu.vue';
import Cart from './components/Cart.vue';

export default {
  components: {
    NavBar,
    Menu,
    Cart
  },
  data() {
    return {
      cartItems: []
    };
  },
  methods: {
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