<template>
    <div class="wrapper">
        <menu-item v-for="item in items" :key="item.id" :item="item" @add-to-cart="addToCartHandler(item)"></menu-item>
    </div>

    <div class="clearfix"></div>
</template>

<script>
import MenuItem from './MenuItem.vue'

  export default {
    components: {
      MenuItem
    },
    data() {
      return {
        items: [
          {
            name: 'Carne Guisada',
            price: 12.00
          },
          {
            name: 'Tamales',
            price: 15.00
          }
      ]}
    },
    methods: {
      addToCartHandler(item) {
        this.$emit('add-to-cart', item);
      }
    }
  }
  
</script>

<style scoped>
.center {
  margin: auto;
  width: 100%;
  border: 3px solid green;
  padding: 10px;
}

.wrapper {
    margin: 2rem;
    display: grid;
    margin-left: 10vw;
    margin-right: 10vw;
    grid-gap: 5vw;
    background-color: #fff;
    color: #444;
}

@media (min-width: 200px) {
  .wrapper {
    margin-left: 12vw;
    margin-right: 12vw;
    grid-template-columns: 33vw 33vw; /* controls the number of columns */
  }
}

@media (min-width: 600px) {
  .wrapper {
    grid-template-columns: 25vw 25vw 25vw; /* controls the number of columns */
  }
}

@media (min-width: 1024px) {
  .wrapper {
    margin-left: 8vw;
    margin-right: 8vw;
    grid-template-columns: 16vw 16vw 16vw 16vw; /* controls the number of columns */
  }
}
</style>
