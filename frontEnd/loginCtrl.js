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
            window.sessionStorage.token = response.data
        }, function myError(response) {
            console.log("Failed to login $http.$$url")
        });
    }
});