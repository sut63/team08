/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDiagnose,
    EntDiagnoseFromJSON,
    EntDiagnoseFromJSONTyped,
    EntDiagnoseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDepartmentEdges
 */
export interface EntDepartmentEdges {
    /**
     * DepartmentDiagnose holds the value of the department_diagnose edge.
     * @type {Array<EntDiagnose>}
     * @memberof EntDepartmentEdges
     */
    departmentDiagnose?: Array<EntDiagnose>;
}

export function EntDepartmentEdgesFromJSON(json: any): EntDepartmentEdges {
    return EntDepartmentEdgesFromJSONTyped(json, false);
}

export function EntDepartmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDepartmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'departmentDiagnose': !exists(json, 'departmentDiagnose') ? undefined : ((json['departmentDiagnose'] as Array<any>).map(EntDiagnoseFromJSON)),
    };
}

export function EntDepartmentEdgesToJSON(value?: EntDepartmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'departmentDiagnose': value.departmentDiagnose === undefined ? undefined : ((value.departmentDiagnose as Array<any>).map(EntDiagnoseToJSON)),
    };
}

