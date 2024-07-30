<template>
  <div>
    <div id="checkout">
    <!-- Checkout will insert the payment form here -->
    </div>
  </div>
</template>

<script>
export default {

};
</script>

<style scoped>
.payment-form {
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #f9f9f9;
}

.payment-form h2 {
  margin-bottom: 20px;
  font-size: 24px;
  text-align: center;
  color: #333;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 12px;
  border: 1px solid #ced4da;
  border-radius: 4px;
}

.card-row {
  margin-bottom: 20px;
}

.card-element {
  margin-bottom: 10px;
  font-size: 10rem;
}

.card-number,
.card-expiry-cvc {
  width: 100%;
}

.pay-button {
  width: 100%;
  padding: 12px;
  background-color: #6772e5;
  border: none;
  border-radius: 4px;
  color: #fff;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.pay-button:hover {
  background-color: #5469d4;
}

.error-message {
  color: #fa755a;
  font-size: 14px;
  text-align: center;
}
</style>
