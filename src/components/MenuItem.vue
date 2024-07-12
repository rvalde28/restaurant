<template>
    <div>
      <div class="image-div">
        <img src="../assets/images/YH_May23_Chicken_Nachos.jpg" ref="nachos">
      </div>
      <div>
          <div class="item-name">
              {{ item.name }}
          </div>

          <div class="item-price">
              ${{ item.price }}
          </div >

          <div class="item-count center-div">
            <div class="input-count-grid">
              <button class="count-button" @click="countSubtract()"> - </button>
              <input class="count-number" type="number"  v-model="count">
              <button class="count-button" @click="countAdd()"> + </button>
            </div>
          </div >

          <div class="cart-button-div">
            <button class="add-cart"  @click="addToCart(item)"> Add to Cart </button>
          </div >
      </div>
    </div>
</template>

<script>
  const maxOrders = 25
  export default {
    props: ['item'],
    data() {
    return {
      count: 0,
    };
  },
    methods: {
      addToCart(item) {
        let itemInstance={
        name: item.name,
        price: item.price,
        count: this.count,
        }
        console.log("called")
        this.$emit('add-to-cart', itemInstance);
      },
      countAdd(){
        if (this.count >= maxOrders){
          this.count = maxOrders
          return
        }
        this.count++
      },
      countSubtract(){
        if (this.count <= 0){
          this.count=0
          return
        } else{
          this.count--
        }
      }
    }
  };
</script>
  
<style scoped>
.item-count{
  font-size: .9rem;
  text-align: center;
}

input{
  border:0;
  text-align: center;
}

.center-text {
  margin: auto;
  width: 50%;
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
  margin-bottom: 1vh;
}

.input-count-grid{
  display: grid;
  grid-template-columns: 4vw 10vw 4vw;
  background-color: white;
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
  font-size:1.2vw;
  font-weight: 550 !important;
  border:0;
  cursor:pointer;
  margin:1em;
}

.item-name{
  margin-bottom: .3vh;
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
  position: relative;
  width: 100%;
  object-fit: contain;
  margin-left: auto;
  margin-right: auto;
}

@media (min-width: 360px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 4vw 12vw 4vw;
  }

  .item-name{
    font-size: 5vw;
  }
  .item-price{
    font-size: 4vw;
  }
  .count-button{
    font-size: 4vw;
    font-weight: lighter;
  }
  .count-number{
    font-size: 3.5vw;
  }
  .add-cart{
    font-size: 4.4vw;
  }
}

@media (min-width: 512px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 4vw 12vw 4vw;
  }

  .item-name{
    font-size: 5vw;
  }
  .item-price{
    font-size: 4vw;
  }
  .count-button{
    font-size: 4vw;
    font-weight: lighter;
  }
  .count-number{
    font-size: 3.5vw;
  }
  .add-cart{
    font-size: 4.4vw;
  }
}

@media (min-width: 720px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 2vw 5vw 2vw;
  }

  .item-name{
    font-size: 2.2vw;
  }
  .item-price{
    font-size: 2.3vw;
  }
  .count-button{
    font-size: 2.5vw;
    font-weight: lighter;
  }
  .count-number{
    font-size: 2.5vw;
  }
  .add-cart{
    font-size: 2.4vw;
  }
}
@media (min-width: 1080px) {
  .input-count-grid {
    display: grid;
    grid-template-columns: 2vw 5vw 2vw;
  }

  .image-div{
    padding: .7vw;
  }

  .item-name{
    font-size: 1.2vw;
  }
  .item-price{
    font-size: 1.1vw;
  }
  .count-button{
    font-size: 1.2vw;
    font-weight: lighter;
  }
  .count-number{
    font-size: 1.1vw;
  }
  .add-cart{
    font-size: 1.1vw;
  }
}

</style>
  