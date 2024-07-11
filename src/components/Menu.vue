<template>
  <div class="wrapper center-div">
    <menu-item v-for="item in menuItems" :key="item.id" :item="item" @add-to-cart="addToCartHandler"></menu-item>
  </div>
</template>

<script>
import MenuItem from './MenuItem.vue';

export default {
  components: {
    MenuItem
  },
  data() {
    return {
      menuItems: [
        { id: 1, name: 'Carne Guisada', price: 10 },
        { id: 2, name: 'Tamales', price: 15 },
        { id: 3, name: 'Tacos', price: 7.5 },
        { id: 1, name: 'Carne Guisada', price: 10 },
        { id: 2, name: 'Tamales', price: 15 },
        { id: 3, name: 'Tacos', price: 7.5 },
      ]
    };
  },
  methods: {
    addToCartHandler(item) {
      this.$emit('add-to-cart', item);
    }
  }
};
</script>

<style scoped>
.center {
  margin: auto;
  width: 100%;
  border: 3px solid green;
  padding: 10px;
}

.center-div{
  max-width: fit-content;
  margin-left: auto;
  margin-right: auto;
}

.wrapper {
    position: relative;
    /* margin: 2rem; */
    margin-left: auto;
    margin-right: auto;
    display: grid;
    grid-gap: 5vw;
    background-color: #efece6;
    color: #444;
}

@media (min-width: 360px) {
  .wrapper {
    grid-template-columns: 75vw; /* controls the number of columns */
  }
}

@media (min-width: 512px) {
  .wrapper {
    grid-template-columns: 75vw; /* controls the number of columns */
  }
}

@media (min-width: 720px) {
  .wrapper {
    grid-template-columns: 35vw 35vw; /* controls the number of columns */
  }
}
@media (min-width: 1080px) {
  .wrapper {
    margin-top: 1.5vw;
    top: 2.2vw;;
    grid-gap: 3vw;
    grid-template-columns: 18vw 18vw 18vw ; /* controls the number of columns */
  }

}
</style>
