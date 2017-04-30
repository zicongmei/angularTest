app.controller('experimrntCtrl', function($scope, $location, $http) {
    $scope.click = [];
    $scope.firstPlayer = false;
    $scope.mouseLocation = function(myE) {
        $scope.x = myE.clientX;
        $scope.y = myE.clientY;
    }
    $scope.mouseClick = function(row, col) {
        $http({
            method: "PUT",
            url: "request/move?row=" + row + "&col=" + col,
            headers: {
                'Authorization': window.sessionStorage.token,
                'Accept': 'application/json',
            }
        }).then(function mySuccess(response) {
            var loc = {
                x: row,
                y: col,
                firstPlayer: $scope.firstPlayer
            };
            $scope.click.push(loc)
            $scope.firstPlayer = !$scope.firstPlayer
        }, function myError(response) {
            console.log("Failed to PUT $http.$$url")
        });
    }
});