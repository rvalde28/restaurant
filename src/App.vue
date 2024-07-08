<script setup>
import NavBar from './components/NavBar.vue'
import Menu from './components/Menu.vue'
// import Cart from './components/Cart.vue';
</script>

<template>
  <!-- <Cart ref="cart"  /> -->

  <div class="top-nav">
    <NavBar/>
  </div>
  <div class="menu">
    <Menu @add-to-cart="addToCart" />
  </div>
  <Cart :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" />

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

<style scoped>
.menu {
  position:relative;
  top: 10vh;
  left:0px;
  z-index: -1;
  /* width: 100vw; */
}

.top-nav{
  height: 10vh;
  position: fixed;
  top: 0px;
  left: 0px;
  width: 100vw;
}

header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;

}

@media (min-width: 200px) {
  .top-nav {
    height: 10vh;
  }
}


@media (min-width: 600px) {
  .top-nav {

  }
}

@media (min-width: 1024px) {
  .top-nav {

  }
}

/* div {
  display: inline-block;
} */
</style>
