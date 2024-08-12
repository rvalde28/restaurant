<template>
  <!-- Desktop Nav -->
  <nav class="bg-white hidden md:flex justify-between items-center fixed top-0 w-full z-10 shadow-md">
    <!-- Left-aligned link -->
    <div>
      <a href="/" @click="navigateHome"
        class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Home</a>


    </div>

    <!-- Right-aligned links -->
    <div class="flex items-center">
      <a href="#"
        class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link
        1</a>
      <a href="#"
        class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link
        2</a>
      <a href="#"
        class="text-gray-800 px-3 py-2 text-lg hover:bg-customBlue hover:text-white transition-colors duration-300 flex items-center">Link
        3</a>

      <button @click="toggleCart"
        class="bg-blue-500 text-white px-3 py-2 rounded-full hover:bg-blue-400 flex items-center mr-4 relative">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-1" fill="none" viewBox="0 0 24 24"
          stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="M15.75 10.5V6a3.75 3.75 0 1 0-7.5 0v4.5m11.356-1.993 1.263 12c.07.665-.45 1.243-1.119 1.243H4.25a1.125 1.125 0 0 1-1.12-1.243l1.264-12A1.125 1.125 0 0 1 5.513 7.5h12.974c.576 0 1.059.435 1.119 1.007ZM8.625 10.5a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Zm7.5 0a.375.375 0 1 1-.75 0 .375.375 0 0 1 .75 0Z" />
        </svg>
        <div>{{ totalItems }}</div>
      </button>

      <!-- Include Cart component -->
      <Cart @remove-item="removeItem" id="mySideNav" @clear-cart="clearCart" @toggle-cart="toggleCart"
        :cartItems="cartItems" :show-cart="showCart" :is-open="isCartOpen" />

    </div>
  </nav>

  <!-- Mobile Nav -->
  <nav class="btm-nav md:hidden z-[1000] rounded-t-3xl bg-blue-500">
    <a href="/" @click="navigateHome" class="text-white text-xl">
      <span class="active-mbl"><AnOutlinedHome /></span>
    </a>

    <a class="text-white text-xl">
      <span class=""><IcProfileCircle /></span>
    </a>

    <button @click="toggleCart" class="text-white indicator">
      <!-- add active-mbl to div if active -->
      <div class="relative text-xl">
        <AkShoppingBag />
        <span v-if="cartItems.length > 0"
          class="indicator-item badge bg-white text-blue-500 absolute top-0 right-0 p-1.5">{{ totalItems }}</span>
      </div>
    </button>
  </nav>
</template>

<script>
import { IcProfileCircle } from '@kalimahapps/vue-icons';
import { AkShoppingBag } from '@kalimahapps/vue-icons';
import { AnOutlinedHome } from '@kalimahapps/vue-icons';
import Cart from './Cart.vue';
import { globalState, getTotalCount } from '../globalstate';
import { computed } from 'vue';

// let totalItems = globalState.totalItems
const totalItems = computed(() => globalState.totalItems);

export default {
  props: ["showCart"],
  components: {
    Cart,
    IcProfileCircle,
    AkShoppingBag,
    AnOutlinedHome
  },
  setup() {
    let cartItems = globalState.cartItems;

    console.log("rendering cartItems: ", cartItems)

    const removeItem = (index) => {
      cartItems.splice(index, 1); 
      localStorage.setItem('cartItems', JSON.stringify(cartItems));
      getTotalCount();
    };

    const clearCart = () => {
      cartItems = []
      getTotalCount();
    };

    return {
      cartItems,
      clearCart,
      removeItem
    };
  },
  data() {
    return {
      useMobileNav: false,
      windowWidth: window.innerWidth,
      isCartOpen: false,
    }
  },
  methods: {
    navigateHome() {
      this.$router.push('/');
    },
    onResize() {
      this.windowWidth = window.innerWidth
    },
    toggleCart() {
      this.isCartOpen = false
      this.$emit('expand-cart')
    }
  },
  mounted() {
    this.$nextTick(() => {
      if (this.windowWidth < 800) {
        this.useMobileNav = true
      } else {
        this.useMobileNav = false
      }
      window.addEventListener('resize', this.onResize);
    })
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.onResize);
  },
};
</script>

<style scoped>
/* Custom hover effect for links */
a:hover {
  background-color: #57B8FF;
  color: #ffffff;
  /* Set text color to white on hover */
}

/* Ensure the link takes up the full height */
a {
  display: block;
  height: 100%;
}

/* Optionally, adjust text and link alignment */
a,
a:hover {
  text-decoration: none;
  /* Remove underline on hover */
  display: flex;
  align-items: center;
}

.active-mbl {
  background: rgba(255, 255, 255, 1);
  color: rgb(59 130 246 / var(--tw-bg-opacity));
  padding: .25rem;
  border-radius: .5rem;
}
</style>
