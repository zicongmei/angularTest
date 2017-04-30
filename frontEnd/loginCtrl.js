app.controller('loginCtrl', function($scope, $http) {
    $scope.user = "unknown";
    $scope.guestLogin = function(name) {
        console.log("login:" + name)
        $scope.user = name;
        $http({
            method: "POST",
            url: "authenticate/login",
            data: {
                "user": name
            },
        }).then(function mySucces(response) {
            console.log(response)
        }, function myError(response) {
            console.log("Failed to login $http.$$url")
        });
    }
});