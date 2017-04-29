app.controller('loginCtrl', function($scope, $http) {
    $scope.user = "unknown";
    $scope.guestLogin = function(name) {
        console.log("login:" + name)
        $scope.user = name;
    }
});