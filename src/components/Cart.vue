<template>
  <div>
    <div v-if="isOpen || showCart" id="cartBackground" @click="closePopupOnOutsideClick" class="fixed top-0 right-0 bottom-0 left-0 z-10 bg-gray-800 bg-opacity-0"></div>
    <transition name="slide-down">
      <div  v-if="isOpen || showCart" class="z-[11] bg-white w-1/5 border rounded-lg shadow-lg absolute mt-2 top-full right-5 w-128">
        <div class="close-link fixed text-xl pl-2">
          <div class="pt-2">
            <a class="text-secondary hover:bg-third hover:rounded-full py-1 px-2" href="javascript:void(0)" @click="closeNav()">&times;</a>
          </div>
        </div>

        <!-- Cart content here -->
        <div class="py-4 px-2">
          <h3 class="text-lg font-semibold mb-4 text-center">Your Cart</h3>
          <!-- Display cart items here -->
          <ul class="pb-4">
            <li class="py-2 flex " v-for="(item, index) in cartItems" :key="index">
              <div class="flex group hover:bg-third hover:rounded-lg p-3 w-4/5">
                <div>
                  ({{ item.count }}) {{ item.name }} - ${{ item.price * item.count }}
                </div>
                <div id="editItem" class="pl-3 text-white group-hover:text-black">
                  Edit
                </div>
              </div>
              <div @click="removeItem(index)" class="hover:bg-fourth hover:rounded-lg pl-2 pr-2  flex justify-between items-center">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5">
                  <path stroke-linecap="round" stroke-linejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                </svg>
              </div>
            </li>
          </ul>
          <!-- <div class="mt-4">
            Total: ${{ getTotalPrice() }}
          </div> -->

          <!-- <button @click="clearCart" class="bg-blue-500 text-white px-3 py-2 rounded-full hover:bg-blue-400 flex items-center mb-4 relative">Clear Cart</button> -->
          <button @click="navigateCheckout" class="bg-blue-500 w-11/12 text-white ml-4 mr-4 px-3 py-2 rounded-full hover:bg-blue-400 flex items-center mb-4 relative">
            <div class="pl-3 text-lg"> <!-- Use flex-grow to make this div take remaining space -->
              Checkout
            </div>
            <div class="absolute right-0 pr-3 text-lg"> <!-- Use absolute positioning and right-0 to align to the right -->
              ${{ getTotalPrice() }}
            </div>
          </button>
        </div>
      </div>
    </transition>
    
  </div>
</template>

<script>
export default {
  props: ['cartItems', 'isOpen', 'showCart'],
  methods: {
    navigateCheckout() {
      this.$router.push('/checkout');
      this.$emit('toggleCart')
    },
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
      // return this.cartItems.reduce((acc, item) => acc + item.price, 0);
      let total = 0.0
      this.cartItems.forEach(item => {
        total += item.price * item.count;
      });

      return total.toFixed(2)
    }
  }
};
</script>

<style scoped>

</style>
