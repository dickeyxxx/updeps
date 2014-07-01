angular.module('app')
.service('PackageSvc', function ($http) {
  this.fetch = function () {
    return $http.get('/api/packages')
  }
  this.create = function (pkg) {
    return $http.post('/api/packages', pkg)
  }
  this.refresh = function (pkg) {
    return $http.post('/api/packages/refresh')
  }
  this.find = function (path) {
    return $http.get('/api/packages/' + path)
  }
})
