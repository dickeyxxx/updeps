angular.module('app')
.controller('AdminCtrl', function ($scope, CategorySvc) {
  $scope.addCategory = function (name) {
    $scope.category = null
    CategorySvc.create(name)
    reload()
  }

  function reload() {
    CategorySvc.fetch().success(function (categories) {
      $scope.categories = categories
    })
  }

  reload()
})
