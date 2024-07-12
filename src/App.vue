<template>
  <div>
    <NavBar :cartItems="cartItems" />
    <MenuComponent class="mt-14" @add-to-cart="addToCart" />
    <Cart class="cart-bar" :cart-items="cartItems" @remove-item="removeItem" @clear-cart="clearCart" />

    <div class="fixed bottom-4 right-4">
      <button @click="expandCart()" class="flex items-center justify-center w-12 h-12 bg-blue-500 text-white rounded-full shadow-md">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 10.5V6a3.75 3.75 0 1 0-7.5 0v4.5m11.356-1.993 1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 0 1-1.12-1.243l1.264-12A1.125 1.125 0 0 1 5.513 7.5h12.974c.576 0 1.059.435 1.119 1.007ZM8.625 10.5a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm7.5 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z" />
        </svg>

        <span class="ml-1">{{ cartItems.length }}</span>
      </button>
    </div>
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