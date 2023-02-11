<template>
  <div class="logo overflow-hidden w-full flex items-center" :style="{height:store.headerHeight+'px'}">
    <img class="logo_icon block m-auto" v-if="store.siderCollapsed" width="30" height="30" src="@/assets/logo_small.png" alt="moss" />
    <template v-else>
      <div class="w-full text-center logo_text">MOSS</div>
    </template>
  </div>
</template>

<script setup>
  import {useStore} from "@/store";

  const store = useStore()
</script>


<style scoped>
@font-face {
  font-family: "myFont";
  src: url("@/assets/font.otf");
  font-weight: 400;
}
.logo {
  will-change: filter;
}
.logo_icon {
  transition: all .2s;
  filter: drop-shadow(0 0 .5em #64ffb8aa);
}
.logo:hover .logo_icon {
  filter: drop-shadow(0 0 .6em rgba(126, 252, 194));
}

.logo_text{
  font-family: 'myFont',serif;
  font-size: 26px;
  background-image: linear-gradient(315deg, #647eff 30%, #00d734);
  -webkit-background-clip: text;
  color: transparent;
}
</style>