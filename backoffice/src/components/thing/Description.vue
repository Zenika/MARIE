<template>
  <div class="subheading">
    Id : {{thing.id}}<br />
    MacAddress : {{thing.macaddress}}<br />
    IPAddress : {{thing.ipaddress}}<br />
    Type : {{thing.type}}<br />
    Protocol : {{thing.protocol}}<br />
    Location : {{thing.location}}
  </div>
</template>
<script>
export default {
  name: 'marie-thing-description',
  props: ['thing']
}
</script>
