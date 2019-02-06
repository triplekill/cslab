package org.exfly.java8;

import org.junit.Test;

import java.util.Arrays;

@FunctionalInterface
interface FunctionalDefaultMethods {
    void method();

    default void defaultMethod() {
    }
}

/**
 *  https://www.jianshu.com/p/5b800057f2d8
 */
public class Characteristic {
    @Test
    public void testLambda() {
        Arrays.asList( "a", "b", "d" ).forEach(e -> System.out.println( e ) );
        Arrays.asList( "a", "b", "d" ).forEach( ( String e ) -> System.out.println( e ) );
        Arrays.asList( "a", "b", "d" ).forEach( e -> {
            System.out.print( e );
            System.out.print( e );
        } );

        //  Lambda表达式可以引用类成员和局部变量（会将这些变量隐式得转换成final的）
        String separator = ",";
        Arrays.asList( "a", "b", "d" ).forEach(
                ( String e ) -> System.out.print( e + separator ) );
//        final String separator = ",";
//        Arrays.asList( "a", "b", "d" ).forEach(
//                ( String e ) -> System.out.print( e + separator ) );

        Arrays.asList( "a", "b", "d" ).sort( ( e1, e2 ) -> e1.compareTo( e2 ) );
        Arrays.asList( "a", "b", "d" ).sort( ( e1, e2 ) -> {
            int result = e1.compareTo( e2 );
            return result;
        } );



    }
}
