<template>
  <div>
    <div v-if="isOpen" id="cartBackground" @click="closePopupOnOutsideClick" class="fixed top-0 right-0 bottom-0 left-0 z-10 bg-gray-800 bg-opacity-0"></div>
    <transition name="slide-down">
      <div  v-if="isOpen" class="z-[11] bg-white border rounded-lg shadow-lg absolute mt-2 top-full right-0 w-64">
        <div class="close-link">
          <a href="javascript:void(0)" @click="closeNav()">&times;</a>
        </div>  


        <!-- Cart content here -->
        <div class="p-4">
          <h3 class="text-lg font-semibold mb-4">Your Cart</h3>
          <!-- Display cart items here -->
          <ul>
            <li v-for="(item, index) in cartItems" :key="index">
              {{ item.name }} - ${{ item.price }}
            </li>
          </ul>
          <!-- Example of total price calculation -->
          <div class="mt-4">
            Total: ${{ getTotalPrice() }}
          </div>
          <button @click="clearCart">Clear Cart</button>
        </div>
      </div>
    </transition>
    
  </div>
</template>

<script>
export default {
  props: ['cartItems', 'isOpen'],
  methods: {
    closePopupOnOutsideClick(event) {
      console.log("outside click")
       // Check if the click target is the background overlay or element with id "windowBackground"
      if (event.target.id === 'cartBackground') {
        this.$emit('toggleCart')
      } 
    },
    closeNav(){
      this.$emit('toggleCart')
    },
    removeItem(index) {
      this.$emit('remove-item', index); // Emit event to remove item from cart
    },
    clearCart() {
      this.$emit('clear-cart'); // Emit event to clear the cart
    },
    getTotalPrice() {
      return this.cartItems.reduce((acc, item) => acc + item.price, 0);
    }
  }
};
</script>

<style scoped>

</style>
