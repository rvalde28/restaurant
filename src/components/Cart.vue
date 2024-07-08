<template>
    <div>
      <h2>Shopping Cart</h2>
      <ul>
        <li v-for="(item, index) in cartItems" :key="index">
          {{ item.name }} - ${{ item.price }}
          <button @click="removeItem(index)">Remove</button>
        </li>
      </ul>
      <p>Total: ${{ getTotal() }}</p>
      <button @click="clearCart">Clear Cart</button>
    </div>
  </template>
  
  <script>
  export default {
    props: ['cartItems'],
    methods: {
      removeItem(index) {
        this.$emit('remove-item', index); // Emit event to remove item from cart
      },
      getTotal() {
        return this.cartItems.reduce((total, item) => total + item.price, 0).toFixed(2);
      },
      clearCart() {
        this.$emit('clear-cart'); // Emit event to clear the cart
      }
    }
  };
  </script>
  