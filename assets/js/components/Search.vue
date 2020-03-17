<template>
  <div class="search">
    <div class="container">
      <img
        class="search-logo"
        src="http://localhost:8080/assets/img/search-gopher.png/"
        width="400px"
      />
      <b-form-input
        class="input-search"
        v-model="query"
        placeholder="Please input search key word."
        style="width: 500px"
      ></b-form-input>
      <b-button class="button-search" variant="outline-dark" v-on:click="getResult(query)">Search</b-button>

      <div class="book-wrapper">
        <div class="book-card" v-for="item in items" :key="item.volumeInfo">
          <b-card no-body class="overflow-hidden" style="max-width: 540px">
            <b-row no-gutters>
              <b-col md="6">
                <b-card-img v-bind:src="item.volumeInfo.imageLinks.thumbnail"></b-card-img>
              </b-col>
              <b-col md="6">
                <b-card-body
                  v-bind:title="item.volumeInfo.title"
                  v-bind:sub-title="item.volumeInfo.subtitle"
                >
                  <b-card-text>
                    {{ item.volumeInfo.authors }}
                    <br />
                    {{ item.volumeInfo.publisher }}
                    <br />
                    {{ item.volumeInfo.publishedDate }}
                  </b-card-text>
                  <b-button
                    variant="outline-dark"
                    v-bind:href="item.volumeInfo.previewLink"
                    target="_blank"
                  >Show Detail</b-button>
                  <br />
                  <b-button class="button-add" variant="outline-dark" target="_blank">Add</b-button>
                </b-card-body>
              </b-col>
            </b-row>
          </b-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
module.exports = {
  data() {
    return {
      query: "",
      items: []
    };
  },
  methods: {
    getResult(query) {
      axios
        .get("https://www.googleapis.com/books/v1/volumes?q=search" + query)
        .then(response => {
          console.log(response.data);
          this.items = response.data.items;
        });
    }
  }
};
</script>

<style>
.search-logo {
  display: block;
  margin: 0 auto;
}

.input-search {
  text-align: center;
  display: block;
  margin: 0 auto;
  margin-top: 10px;
}

.button-search {
  display: block;
  margin: 0 auto;
  margin-top: 10px;
}

.book-wrapper {
  width: 100%;
}
</style>