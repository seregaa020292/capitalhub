<template>
  <div class="info-list">
    <ul class="custom-marcker-one">
      <li>Акции</li>
      <li>Облигации</li>
      <li>Фонды ETF</li>
      <li>Криптовалюты</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'InfoList',
}
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';

.info-list {
  margin-top: 50px;
  z-index: -2;
  background: linear-gradient(110deg, rgba(37, 46, 53, 0.5) -13.2%, rgba(37, 46, 53, 0) 123.22%);
  border-radius: 36px;
  padding: 30px;

  @include media-point('sm') {
    width: calc(100% - 70px);
  }

  ul li {
    font-style: normal;
    font-weight: normal;
    font-size: 22px;
  }

  ul li:not(:last-child) {
    margin-bottom: 25px;
  }
}

.custom-marcker-one > li {
  padding: 0 0 0 52px;
  position: relative;
}

.custom-marcker-one > li::before {
  content: '';
  top: calc(50% - 15.5px);
  width: 31px;
  position: absolute;
  left: 0;
  height: 31px;
  background-image: url("data:image/svg+xml,%3Csvg width='31' height='31' viewBox='0 0 31 31' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Crect x='0.5' y='0.5' width='30' height='30' rx='15' fill='url(%23paint0_linear)' fill-opacity='0.5' stroke='url(%23paint1_linear)'/%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M14.1605 20.0118C14.0424 19.9819 13.9306 19.9207 13.8382 19.8283L11.3383 17.3283C11.065 17.055 11.065 16.6117 11.3383 16.3384C11.6117 16.065 12.0549 16.065 12.3283 16.3384L14.3333 18.3435L19.6717 13.0051C19.9451 12.7317 20.3883 12.7317 20.6617 13.0051C20.935 13.2784 20.935 13.7217 20.6617 13.995L14.8283 19.8284C14.6473 20.0094 14.3918 20.0705 14.1605 20.0118Z' fill='url(%23paint2_linear)'/%3E%3Cdefs%3E%3ClinearGradient id='paint0_linear' x1='25.8333' y1='37.975' x2='1.80833' y2='1.14849e-06' gradientUnits='userSpaceOnUse'%3E%3Cstop stop-color='%23272E37'/%3E%3Cstop offset='1' stop-color='%23272E37' stop-opacity='0'/%3E%3C/linearGradient%3E%3ClinearGradient id='paint1_linear' x1='15.5' y1='0' x2='15.5' y2='31' gradientUnits='userSpaceOnUse'%3E%3Cstop stop-color='%23313E47'/%3E%3Cstop offset='1' stop-color='%23313E47' stop-opacity='0'/%3E%3C/linearGradient%3E%3ClinearGradient id='paint2_linear' x1='16' y1='12.8' x2='16' y2='20.0334' gradientUnits='userSpaceOnUse'%3E%3Cstop stop-color='%23309FF6'/%3E%3Cstop offset='1' stop-color='%233DBDFF'/%3E%3C/linearGradient%3E%3C/defs%3E%3C/svg%3E%0A");
  background-repeat: no-repeat;
  background-position: center;
}
</style>
