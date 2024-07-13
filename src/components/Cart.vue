<template>

    <div  class="sidenav">
      <div id="cartBackground" @click="closePopupOnOutsideClick" class="cart-back fixed inset-0 bg-gray-800 opacity-50"></div>

      <div class="close-link">
        <a href="javascript:void(0)" @click="closeNav()">&times;</a>
      </div>      
      <h3>Shopping Cart</h3>
      <div v-for="(item, index) in cartItems" :key="index">
          <div>
            {{ item.name }}
          </div>

          <div>
            ({{ item.count }}) x $ {{ item.price }} =  ${{ item.count * item.price}}
          </div>
          
          <button @click="removeItem(index)">Remove</button>
      </div>

      <!-- <ul>
        <li v-for="(item, index) in cartItems" :key="index">
          {{ item.name }} - ${{ item.price }} -  quantity ({{ item.count }}) - Price ${{ item.count * item.price}}
          <button @click="removeItem(index)">Remove</button>
        </li> 
      </ul>-->

      <p>Total: ${{ getTotal() }}</p>
      <button @click="clearCart">Clear Cart</button>
    </div>
</template>

<script>
export default {
  props: ['cartItems',],
  methods: {
    closePopupOnOutsideClick(event) {
      console.log("outside click")
       // Check if the click target is the background overlay or element with id "windowBackground"
       if (event.target.id === 'cartBackground') {
        this.$emit('expand-cart')
      } 
    },
    closeNav(){
      document.getElementById("mySideNav").style.width="0";
      this.$emit('expand-cart')
    },
    removeItem(index) {
      this.$emit('remove-item', index); // Emit event to remove item from cart
    },
    getTotal() {
      return this.cartItems.reduce((total, item) => total + (item.price * item.count), 0).toFixed(2);
    },
    clearCart() {
      this.$emit('clear-cart'); // Emit event to clear the cart
    }
  }
};
</script>
  
<style>
.cart-back{
  display:none;
  z-index: -1;
}

.close-link{
  padding-top: 15px;
}

h3{
  margin:0;
}

.sidenav{
  bottom:0;
  justify-content: center;
  color:white;
  height: 100%;
  width:0;
  position: fixed;
  z-index: 1;
  top: 0;
  left: 0;
  background-color: #111;
  overflow: hidden;
  overflow-y:auto;
  transition: 0.5s;
}
</style>