/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api;

import io.lakefs.clients.api.ApiException;
import io.lakefs.clients.api.model.CredentialsWithSecret;
import io.lakefs.clients.api.model.Error;
import io.lakefs.clients.api.model.Setup;
import io.lakefs.clients.api.model.StorageConfig;
import io.lakefs.clients.api.model.VersionConfig;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ConfigApi
 */
@Ignore
public class ConfigApiTest {

    private final ConfigApi api = new ConfigApi();

    
    /**
     * 
     *
     * get version of lakeFS server
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getLakeFSVersionTest() throws ApiException {
        VersionConfig response = api.getLakeFSVersion();

        // TODO: test validations
    }
    
    /**
     * 
     *
     * retrieve lakeFS storage configuration
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageConfigTest() throws ApiException {
        StorageConfig response = api.getStorageConfig();

        // TODO: test validations
    }
    
    /**
     * setup lakeFS and create a first user
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void setupTest() throws ApiException {
        Setup setup = null;
        CredentialsWithSecret response = api.setup(setup);

        // TODO: test validations
    }
    
}
