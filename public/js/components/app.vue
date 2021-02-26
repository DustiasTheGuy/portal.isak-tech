<template>
    <div class="image-container">
        <img 
        v-for="(s, index) of sites" 
        v-bind:key="index" 
        :src="s.imgUrl" 
        :title="s.description"
        @click="navigate(s.url)">
    </div>
</template>

<script> 
export default {
    name: "AppComponent",
    data() {
        return {
            sites: [],
            production: true
        }
    },
    created() {
        fetch(this.production ? 'https://portal.isak-tech.tk/data' : 'http://localhost:8083/data')
        .then(response => response.json())
        .then(data => this.sites = data.data);
    },
    methods: {
        navigate: (url) => window.location.href = url
    },
}
</script> 
