/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 * 
 *   http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

var TrafficPortalService = function($http, $q, $window, propertiesModel) {

    this.getReleaseVersionInfo = function() {
        var deferred = $q.defer();
        $http.get('traffic_portal_release.json')
            .then(
                function(result) {
                    deferred.resolve(result);
                },
                function(fault) {
                    deferred.reject(fault);
                }
            );

        return deferred.promise;
    };

    this.getProperties = function() {
        var deferred = $q.defer();
        $http.get('traffic_portal_properties.json')
            .then(
                function(result) {
                    deferred.resolve(result.data.properties);
                }
            );

        return deferred.promise;
    };

    this.dbDump = function() {
        $window.open(propertiesModel.properties.api.baseUrl + '1.3/dbdump', '_blank');
    };

};

TrafficPortalService.$inject = ['$http', '$q', '$window', 'propertiesModel'];
module.exports = TrafficPortalService;
