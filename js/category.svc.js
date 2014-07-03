angular.module('app')
.service('CategorySvc', function ($http) {
  this.fetch = function () {
    return $http.get('/api/categories')
  }
  this.create = function (name) {
    return $http.post('/api/categories', {name: name})
  }
})
