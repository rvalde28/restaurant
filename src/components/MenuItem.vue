<template>
    <div class="box">
      <img src="../assets/images/YH_May23_Chicken_Nachos.jpg" ref="nachos">

      <div>
          <div class="item-name">
              {{ item.name }}
          </div>

          <div class="item-price">
              {{ item.price }}
          </div >

          <div class="item-count center-div">
            <div class="input-count-grid">
              <button class="count-button" @click="count--"> - </button>
              <!-- <div class="count-number">{{ count }}</div> -->
              <button class="count-button" @click="count++"> + </button>
            </div>
          </div >

          <div class="cart-button-div">
            <button class="add-cart"  @click="addToCart"> Add to Cart </button>
          </div >
      </div>
    </div>
</template>

<script>
  export default {
    props: ['item'],
    methods: {
      addToCart(item) {
        this.$emit('add-to-cart', item);
      }
    }
  }
</script>
  
<style scoped>
.item-count{
  font-size: .9rem;
  text-align: center;
}

.center-div{
  display: flex; 
  flex-direction: row; 
  align-items: center; 
  justify-content: center;

}
.count-button{
  background-color: #57B8FF;
  color: black;
  border:0;
  padding:0;
  cursor:pointer;
}
.count-button:hover{
  background-color: #2a97e7;
}

.count-number{
  background-color: white;
  color: black
}
.item-price{
  margin-inline: auto;
}

.input-count-grid{
  display: grid;
  grid-template-columns: 4vw 10vw 4vw;
}
.cart-button-div{
  align-content: center;
  text-align: center;
}

.add-cart {
  align-content: center;
  text-align: center;
  padding:0.1vw 1vw;
  border-radius: 8px;
  color:#fff;
  background-color:#57B8FF;
  font-size:2.2vw;
  font-weight: 550 !important;
  border:0;
  cursor:pointer;
  margin:1em;
}

.item-name{
  font-size: .9rem;
  text-align: center;
}

.item-price{
  font-size: .9rem;
  text-align: center;
}

.box {
  background-color: #444;
  color: #fff;
  border-radius: 5px;
  font-size: 150%;
}
img {
    width: 90%;
    object-fit: contain;
    padding:10px;
    /* padding-top: 18px;
    padding-left: 18px;
    padding-right: 18px; */
  }

@media (min-width: 600px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 2vw 5vw 2vw;
  }
}

@media (min-width: 1024px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 2vw 5vw 2vw;
  }
  .item-name{
    font-size: 1.1vw;
    text-align: center;
  }
  .item-price{
    font-size: 1.1vw;
    text-align: center;
  }
  .count-number{
    font-size: 1.1vw;
    text-align: center;
  }
  /* .add-cart {
    font-size:1.1vw;
  } */
}
</style>
  